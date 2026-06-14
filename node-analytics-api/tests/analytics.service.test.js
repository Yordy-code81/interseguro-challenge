const { calculateStats, isDiagonal, processMatrices } = require('../src/services/analytics.service');

describe('Analytics Service', () => {

    describe('calculateStats', () => {
        it('should calculate stats correctly for positive numbers', () => {
            const Q = [[1, 2], [3, 4]];
            const R = [[5, 6], [7, 8]];
            const stats = calculateStats([Q, R]);
            
            expect(stats.maxValue).toBe(8);
            expect(stats.minValue).toBe(1);
            expect(stats.totalSum).toBe(36);
            expect(stats.average).toBe(4.5);
        });

        it('should handle negative numbers and zeros correctly', () => {
            const Q = [[-1, 0], [-3, 4]];
            const R = [[0, 0]];
            const stats = calculateStats([Q, R]);
            
            expect(stats.maxValue).toBe(4);
            expect(stats.minValue).toBe(-3);
            expect(stats.totalSum).toBe(0);
            expect(stats.average).toBe(0);
        });
    });

    describe('isDiagonal', () => {
        it('should return true for a diagonal matrix', () => {
            const matrix = [
                [5, 0, 0],
                [0, 2, 0],
                [0, 0, -1]
            ];
            expect(isDiagonal(matrix)).toBe(true);
        });

        it('should return false for a non-diagonal matrix', () => {
            const matrix = [
                [5, 1, 0],
                [0, 2, 0],
                [0, 0, -1]
            ];
            expect(isDiagonal(matrix)).toBe(false);
        });

        it('should return false for a non-square matrix', () => {
            const matrix = [
                [1, 0],
                [0, 1],
                [0, 0]
            ];
            expect(isDiagonal(matrix)).toBe(false);
        });
    });

    describe('processMatrices', () => {
        it('should process both matrices and return consolidated object', () => {
            const Q = [
                [1, 0],
                [0, 1]
            ];
            const R = [
                [2, 3],
                [0, 4]
            ];

            const result = processMatrices(Q, R);
            expect(result.max_value).toBe(4);
            expect(result.min_value).toBe(0);
            expect(result.total_sum).toBe(11);
            expect(result.average).toBe(1.375);
            expect(result.is_diagonal.Q).toBe(true);
            expect(result.is_diagonal.R).toBe(false);
        });
    });
});
