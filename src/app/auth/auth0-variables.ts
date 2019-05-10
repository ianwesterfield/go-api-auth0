interface AuthConfig {
  clientId: string;
  domain: string;
  loginCallbackUrl: string;
  logoutCallbackUrl: string;
}

export const AUTH_CONFIG: AuthConfig = {
  clientId: 'znPufIaMMF07tU2ylq1iU1paZ3pgFsBp',
  domain: 'enquiren.auth0.com',
  loginCallbackUrl: 'http://localhost:4000/callback',
  logoutCallbackUrl: 'http://localhost:4000/'
};
