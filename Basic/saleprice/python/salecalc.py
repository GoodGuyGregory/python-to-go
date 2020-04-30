# Calculates the sales price of an entered item
originalPrice = float(input("enter price of item: "))

discount = int(input("enter percentage off: "))

# Calculate precentage
precentOff = discount / 100

priceAdjustment = originalPrice * precentOff

salePrice = originalPrice - priceAdjustment

print("discounted price: ", "{:.2f}".format(salePrice))
