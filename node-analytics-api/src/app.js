const express = require('express');
const cors = require('cors');
const dotenv = require('dotenv');
const analyticsRoutes = require('./routes/analytics.routes');

// Load environment variables
dotenv.config();

const app = express();

// Global Middleware
app.use(cors());
app.use(express.json());

// Routes
app.use('/api', analyticsRoutes);

// Health check endpoint
app.get('/health', (req, res) => {
    res.status(200).json({ status: 'ok', service: 'node-analytics-api' });
});

// Start Server if not imported by tests
if (require.main === module) {
    const PORT = process.env.PORT || 4000;
    app.listen(PORT, () => {
        console.log(`Node Analytics API listening on port ${PORT}`);
    });
}

module.exports = app;