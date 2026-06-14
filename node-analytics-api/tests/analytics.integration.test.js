const request = require('supertest');
const jwt = require('jsonwebtoken');
const app = require('../src/app');

describe('Analytics API Integration Tests', () => {
    
    let validToken;
    const jwtSecret = process.env.JWT_SECRET || 'default_secret_for_challenge';

    beforeAll(() => {
        validToken = jwt.sign({ user: 'tester' }, jwtSecret, { expiresIn: '1h' });
    });

    it('should return 401 if no token is provided', async () => {
        const response = await request(app)
            .post('/api/analytics')
            .send({ Q: [[1]], R: [[1]] });
        
        expect(response.status).toBe(401);
        expect(response.body.success).toBe(false);
    });

    it('should return 401 for an invalid token', async () => {
        const response = await request(app)
            .post('/api/analytics')
            .set('Authorization', 'Bearer invalid_token_123')
            .send({ Q: [[1]], R: [[1]] });
        
        expect(response.status).toBe(401);
    });

    it('should return 400 if Q or R is missing', async () => {
        const response = await request(app)
            .post('/api/analytics')
            .set('Authorization', `Bearer ${validToken}`)
            .send({ Q: [[1]] }); // Missing R
        
        expect(response.status).toBe(400);
        expect(response.body.success).toBe(false);
        expect(response.body.error).toContain('Both Q and R');
    });

    it('should return analytics stats successfully for valid matrices', async () => {
        const response = await request(app)
            .post('/api/analytics')
            .set('Authorization', `Bearer ${validToken}`)
            .send({ 
                Q: [[1, 0], [0, 1]], 
                R: [[2, 0], [0, 2]] 
            });
        
        expect(response.status).toBe(200);
        expect(response.body.success).toBe(true);
        expect(response.body.data.max_value).toBe(2);
        expect(response.body.data.min_value).toBe(0);
        expect(response.body.data.average).toBe(0.75); // (1+0+0+1+2+0+0+2)/8 = 6/8 = 0.75
        expect(response.body.data.total_sum).toBe(6);
        expect(response.body.data.is_diagonal.Q).toBe(true);
        expect(response.body.data.is_diagonal.R).toBe(true);
    });
});
