export const startWs = () => {
    return new WebSocket(
        `ws://${
            process.env.NODE_ENV === "development"
                ? "localhost:3000"
                : process.env.REACT_APP_HOST
        }/api/chat`
    )
}
