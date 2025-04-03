const environment = import.meta.env.VITE_ENVIRONMENT || 'development';

let config = {
  httpProtocol: 'http',
  wsProtocol: 'ws',
  domain: 'localhost:3000'
};

if (environment === 'local') {
  config = {
    httpProtocol: 'http',
    wsProtocol: 'ws',
    domain: 'localhost:3000'
  };
} else {
  config = {
    httpProtocol: 'https',
    wsProtocol: 'ws',
    domain: 'api-broken-thunder-9310.fly.dev'
  };
}

export const API_URL = `${config.httpProtocol}://${config.domain}`;
export const WS_URL = `${config.wsProtocol}://${config.domain}`;