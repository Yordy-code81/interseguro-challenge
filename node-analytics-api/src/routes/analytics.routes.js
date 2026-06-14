const express = require('express');
const router = express.Router();
const analyticsController = require('../controllers/analytics.controller');
const authMiddleware = require('../middleware/auth.middleware');

// Protect route with JWT middleware
router.post('/analytics', authMiddleware.verifyToken, analyticsController.processAnalytics);

module.exports = router;
