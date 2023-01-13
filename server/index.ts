import express, { Request, Response } from 'express'
import * as auth from './src/controllers/authControllers'
import LoadEnv from './src/config/dotenv'
import http from 'http'
import * as io from "socket.io"


LoadEnv()
const app = express();
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.use((req, res, next) => {
    res.setHeader("Access-Control-Allow-Origin", "*")
    next()
})

const socketServer = http.createServer(app)
const socketIO = new io.Server(socketServer, { cors: { origin: "*" } })

socketIO.on('connection', (socket) => {
    console.log(`⚡: ${socket.id} user just connected!`);

    socket.on('disconnect', () => {
      console.log('🔥: A user disconnected');
    });

    socket.on('chat message', (msg) => {
        console.log('message: ' + msg);
        socket.emit('chat message', msg);
      });
});


let PORT = process.env.PORT || 8080;
let SOCKET_PORT = process.env.SOCKET_PORT || 9090

app.get("/", (req: Request, res: Response) => {
    res.status(200).send("<h1>Home Page</h1>")
})


app.post("/api/sing-up", auth.SingUp)
app.post("/api/login", auth.Login)
app.patch("/api/logout", auth.Logout)


app.listen(PORT, () => {
    console.log(`\n Server Listening On "http://localhost:${PORT}" `);

    socketServer.listen(SOCKET_PORT, () => { console.log(`\n Socket Server Listening on "ws://localhost:${SOCKET_PORT}"`) })
});
