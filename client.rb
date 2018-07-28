require 'aws-sdk'
require 'json'

sqs ||= Aws::SQS::Client.new(
  region: 'us-west-2',
  endpoint: 'http://127.0.0.1:9324',
  profile: 'personal',
  access_key_id: 'dummy',
  secret_access_key: 'dummy'
)

queue = sqs.create_queue(
  queue_name: 'MyGroovyQueue'
)

puts queue.queue_url

queue_url = queue.queue_url

message_body = { id: 2, text: 'test queue message' }.to_json

sqs.send_message(queue_url: queue_url, message_body: message_body)
