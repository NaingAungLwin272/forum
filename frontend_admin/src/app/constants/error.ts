export const customError = (error: any) => {
    const errorMessage = error.error.error;
    const startIndex = errorMessage.lastIndexOf('=') + 2;
    const endIndex = errorMessage.length;
    const extractedMessage = errorMessage.substring(startIndex, endIndex);
    return extractedMessage
}