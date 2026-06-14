const jwt = require('jsonwebtoken');

const verifyToken = (req, res, next) => {
    const authHeader = req.headers['authorization'];
    
    if (!authHeader) {
        return res.status(401).json({
            success: false,
            error: 'No token provided'
        });
    }

    const parts = authHeader.split(' ');
    if (parts.length !== 2 || parts[0] !== 'Bearer') {
        return res.status(401).json({
            success: false,
            error: 'Invalid Authorization header format'
        });
    }

    const token = parts[1];
    const secret = process.env.JWT_SECRET || 'default_secret_for_challenge';

    jwt.verify(token, secret, (err, decoded) => {
        if (err) {
            return res.status(401).json({
                success: false,
                error: 'Invalid or expired token'
            });
        }
        
        req.user = decoded;
        next();
    });
};

module.exports = {
    verifyToken
};
