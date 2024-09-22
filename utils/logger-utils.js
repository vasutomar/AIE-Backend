import pino from "pino";

export function createLogger() {
    const logger = pino();
    return logger;
}