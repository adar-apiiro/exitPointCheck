using System;
using System.Threading.Tasks;
using Adyen;
using Adyen.Model.Checkout;
using Adyen.Service;
using Adyen.Service.Checkout;

namespace AdyenDotNetExample
{
    class Program
    {
        static async Task Main(string[] args)
        {
            // Configure your API key and client
            var config = new Config()
            {
                XApiKey = "Your-Adyen-API-Key"
            };
            var client = new Client(config);
            var checkout = new Checkout(client);

            // Example 1: Retrieve available payment methods
            var paymentMethodsRequest = new PaymentMethodsRequest
            {
                MerchantAccount = "YourMerchantAccount"
            };

            try
            {
                var paymentMethodsResponse = await checkout.PaymentMethodsAsync(paymentMethodsRequest);
                Console.WriteLine("Available Payment Methods: " + paymentMethodsResponse);
            }
            catch (Exception ex)
            {
                Console.WriteLine("Error retrieving payment methods: " + ex.Message);
            }

            // Example 2: Make a payment
            var paymentsRequest = new PaymentRequest
            {
                MerchantAccount = "YourMerchantAccount",
                Amount = new Adyen.Model.Checkout.Amount("EUR", 1000), // 10 EUR
                Reference = "YourOrderReference",
                PaymentMethod = new PaymentMethodDetails
                {
                    Type = "scheme", // Card payments
                    EncryptedCardNumber = "test_4111111111111111",
                    EncryptedExpiryMonth = "test_03",
                    EncryptedExpiryYear = "test_2030",
                    EncryptedSecurityCode = "test_737"
                },
                ReturnUrl = "https://your.return.url"
            };

            try
            {
                var paymentsResponse = await checkout.PaymentsAsync(paymentsRequest);
                Console.WriteLine("Payment Response: " + paymentsResponse);
            }
            catch (Exception ex)
            {
                Console.WriteLine("Error processing payment: " + ex.Message);
            }

            // Example 3: Payment details
            var paymentsDetailsRequest = new PaymentDetailsRequest
            {
                Details = new Adyen.Model.Checkout.DefaultPaymentMethodDetails
                {
                    Type = "scheme"
                },
                PaymentData = "YourPaymentData"
            };

            try
            {
                var paymentsDetailsResponse = await checkout.PaymentsDetailsAsync(paymentsDetailsRequest);
                Console.WriteLine("Payment Details Response: " + paymentsDetailsResponse);
            }
            catch (Exception ex)
            {
                Console.WriteLine("Error retrieving payment details: " + ex.Message);
            }
        }
    }
}
