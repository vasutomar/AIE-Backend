import express from 'express';

import { authenticationRouter } from './routers/authentication-router.js';
import { onboardingRouter } from './routers/onboarding-router.js';
import { profileRouter } from './routers/profile-router.js';

import { getVariable } from './config/getVariables.js';
import { createLogger } from './utils/logger-utils.js';

import bodyParser from 'body-parser';
import cors from 'cors';

const logger = createLogger();

const port = getVariable('PORT') | 3001;
const host = getVariable('HOSTNAME') | '127.0.0.1';

const app = express();
app.set('port', port);
app.use(bodyParser.json());
app.use(cors());

app.use(authenticationRouter);
app.use(onboardingRouter);
app.use(profileRouter);

app.listen(parseInt(port, 10), host, () => {
    logger.info(`AIE Server running on ${host}:${port}`);
});
