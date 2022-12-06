# Sequence Diagram - Use Case

As a breakout room, work together to create a sequence diagram (or multiple if you'd prefer) providing clear visualization for the following workflow:

A customer comes to your website to order a custom, made-to-order widget. The customer provides length and width dimensions, desired material type, and quantity. Your system checks to make sure that the dimensions provided are valid numbers - if invalid, an error message is raised and the customer is allowed to reenter and resubmit. The system also checks current inventory to confirm availability of the requested material type. If the material type is not available, the system will reach out to an external supplier via B2B API to order the material to cover the new order.

With material "in hand", the system will initiate production, passing in requested dimensions, requested material type, and quantity. Once production is complete, the system will initiate invoicing to the customer, process payment, and upon receipt of payment, initiate shipping. On completion, a final "thank you" communication is sent to the customer notifying them that their order is on the way.

If the payment processing fails, the customer is allowed to resubmit once. If resubmission is successful, same flow as outlined above. If resubmission fails again, the order is canceled.
