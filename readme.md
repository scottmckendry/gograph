# Go Graph Helper
A simple example of how to authenticate users interactively (supports MFA) with Azure AD/Microsoft Entra and then call the Microsoft Graph API from a Golang application.

I ran into several problems and lacking documentation when trying to do this myself, so I hope this helps someone else.

### Getting Started
1. Create an Azure AD application following the steps [here](https://docs.microsoft.com/en-us/azure/active-directory/develop/quickstart-register-app). Secrets/certificates are not required for user authentication.
2. Create a `.env` file and place it in the root of the project. The file should contain the following values:
```env
CLIENT_ID=<your client id>
TENANT_ID=<your tenant id>
```
