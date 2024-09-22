import express from 'express';
import { authenticationRouter } from './routers/authentication-router.js';
import { getVariable } from './config/getVariables.js';
import bodyParser from 'body-parser';

const port = getVariable('PORT');
const host = getVariable('HOSTNAME');

const app = express();
app.set('port', port);
app.use(bodyParser.json());

app.use(authenticationRouter);
app.listen(parseInt(port, 10), host, () => {
    console.log(`Server started at http://localhost:${port}`);
});
