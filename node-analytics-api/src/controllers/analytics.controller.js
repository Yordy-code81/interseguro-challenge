const analyticsService = require('../services/analytics.service');

const processAnalytics = (req, res) => {
    try {
        const { Q, R } = req.body;
        
        console.log(`[AnalyticsController] Received request to process matrices. Q size: ${Q ? Q.length : 0}, R size: ${R ? R.length : 0}`);

        if (!Q || !R) {
            return res.status(400).json({
                success: false,
                error: 'Both Q and R matrices must be provided in the request body'
            });
        }

        const statistics = analyticsService.processMatrices(Q, R);

        console.log(`[AnalyticsController] Successfully calculated statistics. Returning results.`);

        return res.status(200).json({
            success: true,
            data: statistics
        });
    } catch (error) {
        console.error(`[AnalyticsController] Error processing matrices: ${error.message}`);
        return res.status(500).json({
            success: false,
            error: error.message || 'Internal server error during analytics processing'
        });
    }
};

module.exports = {
    processAnalytics
};
