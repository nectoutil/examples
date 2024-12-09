import jwt from '@necto/crypto/jwt';

const SECRET_KEY = 'your-secret-key-here';

export class JWTService {
  static generateToken(payload, expiresIn = '24h') {
    try {
      return jwt.sign(payload, SECRET_KEY, { expiresIn });
    } catch (error) {
      throw new Error(`Error generating token: ${error.message}`);
    }
  }

  static verifyToken(token) {
    try {
      return jwt.verify(token, SECRET_KEY);
    } catch (error) {
      throw new Error(`Invalid token: ${error.message}`);
    }
  }
}

// Example with Express middleware
export const authMiddleware = (req, res, next) => {
  try {
    const authHeader = req.headers.authorization;
    if (!authHeader?.startsWith('Bearer ')) {
      throw new Error('No token provided');
    }

    const token = authHeader.split(' ')[1];
    const decoded = JWTService.verifyToken(token);
    
    req.user = decoded;
    next();
  } catch (error) {
    res.status(401).json({ error: error.message });
  }
};
