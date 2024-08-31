import 'dotenv/config';

export function getVariable(name) {
    const variableFromEnv = process.env[name];
    return variableFromEnv || '';
}