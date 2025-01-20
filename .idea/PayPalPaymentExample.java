import com.paypal.api.payments;
com.paypal.api.payments.Amount;
com.paypal.api.payments.Details;
import com.paypal.base.rest.APIContext;
com.paypal.api.payments.Transaction;
import com.paypal.api.payments.*;
import com.paypal.base.rest.PayPalRESTException;
import java.util.ArrayList;
import java.util.List;

public class PayPalPaymentExample {

    public static void main(String[] args) {
        String clientId = "YOUR_CLIENT_ID";
        String clientSecret = "YOUR_CLIENT_SECRET";

        APIContext apiContext = new APIContext(clientId, clientSecret, "sandbox");

        Payment payment = new Payment();
        payment.setIntent("sale");

        Payer payer = new Payer();
        payer.setPaymentMethod("paypal");
        payment.setPayer(payer);

        Amount amount = new Amount();
        amount.setCurrency("USD");
        amount.setTotal("10.00"); 

        Transaction transaction = new Transaction();
        transaction.setDescription("Description of the payment transaction.");
        transaction.setAmount(amount);

        List<Transaction> transactions = new ArrayList<>();
        transactions.add(transaction);
        payment.setTransactions(transactions);

        RedirectUrls redirectUrls = new RedirectUrls();
        redirectUrls.setCancelUrl("https://example.com/cancel");
        redirectUrls.setReturnUrl("https://example.com/return");
        payment.setRedirectUrls(redirectUrls);

        try {
            Payment createdPayment = payment.create(apiContext);
            System.out.println("Payment created successfully with ID: " 
                                + createdPayment.getId());
            System.out.println("Payment state: " 
                                + createdPayment.getState());
        } catch (PayPalRESTException e) {
            System.err.println("Error occurred during payment creation:");
            System.err.println(e.getDetails());
        }
    }
}