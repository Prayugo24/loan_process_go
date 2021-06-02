# POST COLLECTION
```bash
    #you can import post collection with this link
    https://www.getpostman.com/collections/e11849c9edff21c704b2

```
# Api Spesification

## Create Loan

Request :
- Method : POST
- Endpoint : `/create_loan`
- Body :
    Form-Data
```json 
{
    params[No_ktp]:999999999
    params[Name]:Andi
    params[Place_of_birth]:Cilacap
    params[Date_of_birth]:1995-08-09
    params[Gender]:Male
    params[Address]:Jakarta
    params[Current_job]:Programmer
    params[Province]: Jawa Timur
    params[No_hp]:09733382
    params[Mothers_name]:Santi
    params[Monthly_income]:3000000 ## minimal 3 jt jika ingin meminjam 1 juta
}
```

Response :

```json 
{
    {
        "response": {
            "Id": 1,
            "Name": "Andi",
            "No_ktp": 999999999,
            "Photo_ktp": "bba534f1-a076-4723-bcc8-dfb4aff693f3.jpeg",
            "Date_of_birth": "1995-08-09",
            "Gender": "Male",
            "Address": "Jakarta",
            "No_hp": 9733382,
            "Mothers_name": "Santi",
            "Status": "Deactived",
            "Total_loan_fund": 1333334,
            "Loan_interest_permonth": 33334,
            "Tenor": {
                "Month_3": 444445,
                "Month_6": 222223,
                "Month_12": 111112
            }
        },
        "status": 200
    }
}
```

## Create Tenor

Request :
- Method : POST
- Endpoint : `/create_tenor`
- Body :

```json 
{
    params[Id_user]:1
    params[Tenor]:3
}
```

Response :

```json 
{
    {
        "response": {
            "Total_loan_fund": 1333334,
            "Handling_fee": 66667,
            "Loans_received": 1266667,
            "Loan_interest_permonth": 33334,
            "Tenor": 3,
            "Monthly_installments": 477779,
            "Id_user_loan": 1,
            "Id": 1
        },
        "status": 200
    }
}
```
## Update Loan

Request :
- Method : POST
- Endpoint : `/update_loan`
- Body :

```json 
{
    params[Id_user]:1
    params[Pay_loan]:477779
    params[Date_pay_loan]:2021-04-29
}
```

Response :

```json 
{
    {
        "response": {
            "id": 1,
            "date": "2021-04-29T00:00:00Z",
            "monthly_bil": 477779,
            "amercement": 0,
            "id_user_loan": 1
        },
        "status": 200
    }
}
```

## Approve Loan

Request :
- Method : POST
- Endpoint : `/approve_loan`
- Body :

```json 
{
    params[Id_user]:1
    params[Date_approve]:2021-04-25
}
```

Response :

```json 
{
    {
        "response": {
            "Name": "Andi",
            "No_ktp": 999999999,
            "Photo_ktp": "bba534f1-a076-4723-bcc8-dfb4aff693f3.jpeg",
            "Date_of_birth": "1995-08-09",
            "Gender": "Male",
            "Address": "Jakarta",
            "Current_job": "Programmer",
            "No_hp": 9733382,
            "Mothers_name": "Santi",
            "Monthly_income": 4000000,
            "Status": "Actived",
            "Loan_interest_permonth": 33334,
            "Total_loan_fund": 1333334,
            "Monthly_installments": 477779,
            "Loans_received": 0,
            "Tenor": 3
        },
        "status": 200
    }
}
```

## Read Loan

Request :
- Method : POST
- Endpoint : `/read_loan`
- Body :

```json 
{
    params[Id_user]:1
}
```

Response :

```json 
{
    {
        "response": {
            "Name": "Andi",
            "No_ktp": 999999999,
            "Photo_ktp": "bba534f1-a076-4723-bcc8-dfb4aff693f3.jpeg",
            "Date_of_birth": "1995-08-09",
            "Gender": "Male",
            "Address": "Jakarta",
            "Current_job": "Programmer",
            "No_hp": 9733382,
            "Mothers_name": "Santi",
            "Monthly_income": 4000000,
            "Status": "Actived",
            "Loan_interest_permonth": 33334,
            "Total_loan_fund": 1333334,
            "Monthly_installments": 477779,
            "Loans_received": 1266667,
            "Tenor": 3
        },
        "status": 200
    }
}
```




