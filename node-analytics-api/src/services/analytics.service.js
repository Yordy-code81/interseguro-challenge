const calculateStats = (matrices) => {
    let max = -Infinity;
    let min = Infinity;
    let sum = 0;
    let count = 0;

    matrices.forEach(matrix => {
        if (!matrix || matrix.length === 0) return;
        
        matrix.forEach(row => {
            row.forEach(val => {
                if (val > max) max = val;
                if (val < min) min = val;
                sum += val;
                count++;
            });
        });
    });

    if (count === 0) {
        return { maxValue: 0, minValue: 0, average: 0, totalSum: 0 };
    }

    // Handle precision issues
    const formatNumber = (num) => Number(num.toFixed(4));

    return {
        maxValue: formatNumber(max),
        minValue: formatNumber(min),
        average: formatNumber(sum / count),
        totalSum: formatNumber(sum)
    };
};

const isDiagonal = (matrix) => {
    if (!matrix || matrix.length === 0 || matrix.length !== matrix[0].length) {
        return false; // Only square matrices can be truly diagonal in the standard sense
    }

    const n = matrix.length;
    for (let i = 0; i < n; i++) {
        for (let j = 0; j < n; j++) {
            // If it's not on the main diagonal and it's not (very close to) zero
            if (i !== j && Math.abs(matrix[i][j]) > 1e-10) {
                return false;
            }
        }
    }
    return true;
};

const processMatrices = (Q, R) => {
    const stats = calculateStats([Q, R]);
    
    const isQDiagonal = isDiagonal(Q);
    const isRDiagonal = isDiagonal(R);

    return {
        max_value: stats.maxValue,
        min_value: stats.minValue,
        average: stats.average,
        total_sum: stats.totalSum,
        is_diagonal: {
            Q: isQDiagonal,
            R: isRDiagonal
        }
    };
};

module.exports = {
    calculateStats,
    isDiagonal,
    processMatrices
};
