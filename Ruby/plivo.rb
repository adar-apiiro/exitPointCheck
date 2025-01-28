require 'plivo'

include Plivo

client = RestClient.new("YourAuthID", "YourAuthToken")

response = client.messages.create(
  'YourSourceNumber',
  ['DestinationNumber'],
  'Hello from Plivo!'
)

puts "Message UUID: #{response.message_uuid}"
