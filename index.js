import express from 'express';
import { authenticationRouter } from './routers/authentication-router.js';
import { onboardingRouter } from './routers/onboarding-router.js';
import { getVariable } from './config/getVariables.js';
import bodyParser from 'body-parser';

import { createLogger } from './utils/logger-utils.js';

const logger = createLogger();

const port = getVariable('PORT') | 3001;
const host = getVariable('HOSTNAME') | '127.0.0.1';

logger.info(`Fetched host : ${host} and port :${port}`);

const app = express();
app.set('port', port);
app.use(bodyParser.json());

app.use(authenticationRouter);
app.use(onboardingRouter);

app.listen(parseInt(port, 10), host, () => {
    logger.info(`AIE Server running on ${host}:${port}`);
});
