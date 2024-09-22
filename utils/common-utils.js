export function handleServerError(err, res) {
    const serverResponse = {
        message: "We hit a snag",
        statusCode: 500,
        data: null,
        error: err.message
    };
    res.send(serverResponse);
}

export function prepareServerResponse(statusCode, message, data) {
    const serverResponse = {
        message,
        statusCode,
        data,
        error: null
    };
    return serverResponse;
}