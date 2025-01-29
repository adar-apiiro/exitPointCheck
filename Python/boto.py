import boto3
import rails
def list_s3_buckets():
    # Initialize an S3 client
    s3 = boto3.client('s3')
    
    # Retrieve a list of S3 buckets
    response = s3.list_buckets()
    
    # Print bucket names
    print("S3 Buckets:")
    for bucket in response['Buckets']:
        print(f"  - {bucket['Name']}")

if __name__ == "__main__":
    print("Listing S3 buckets:")
    list_s3_buckets()
