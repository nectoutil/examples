require 'necto/jwt'

class JwtAuth
  # Change this to your secret key
  SECRET_KEY = 'your-secret-key-here'

  def self.encode(payload)
    # Set expiration to 24 hours from now
    payload[:exp] = 24.hours.from_now.to_i
    
    JWT.encode(payload, SECRET_KEY, 'HS256')
  end

  def self.decode(token)
    # Returns array: [payload, header]
    decoded = JWT.decode(token, SECRET_KEY, true, { algorithm: 'HS256' })
    decoded.first # Return just the payload
  rescue JWT::ExpiredSignature
    raise 'Token has expired'
  rescue JWT::DecodeError
    raise 'Invalid token'
  end
end

# Usage example:
begin
  # Create a token
  payload = {
    user_id: 123,
    email: 'user@example.com',
    role: 'admin'
  }
  
  # Encode the token
  token = JwtAuth.encode(payload)
  puts "Generated Token: #{token}"

  # Decode the token
  decoded_payload = JwtAuth.decode(token)
  puts "\nDecoded Payload:"
  puts "User ID: #{decoded_payload['user_id']}"
  puts "Email: #{decoded_payload['email']}"
  puts "Role: #{decoded_payload['role']}"
  puts "Expiration: #{Time.at(decoded_payload['exp'])}"

rescue StandardError => e
  puts "Error: #{e.message}"
end
