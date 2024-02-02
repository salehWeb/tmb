import { FormEvent, useCallback, useEffect, useState } from 'react'

const socket = new WebSocket("ws://localhost:8080/ws");
const Chat = () => {
    const [massage, setMassage] = useState("");
    const [messages, setMessages] = useState<{ body: string, type: number, clientId: string }[]>([])

    const handelSend = (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        if (!massage) return;
        socket.send(massage);
        setMassage('');
    }

    const handler = useCallback((x: MessageEvent<string>) => {
        const data = JSON.parse(x.data)
        if (data?.body && data?.type) {
            setMessages(prev => [...prev, data])
        }
    }, [])

    useEffect(() => {
        socket.addEventListener("message", handler);
        return () => {
            // eslint-disable-next-line react-hooks/exhaustive-deps
            socket.removeEventListener("message", handler)
        }
    }, [handler])

    return (
        <main>
            <form onSubmit={(e) => handelSend(e)} className="flex flex-col gap-2">
                <label htmlFor="send-massage-input">Send Massage</label>
                <input type="text" id="send-massage-input" value={massage} onChange={(e) => setMassage(e.target.value)} placeholder="..." />
                <button type="submit">Send</button>
            </form>


            <div>
                {messages.map(m => (
                    <div>
                        <p>{m.body}</p>
                        <span>{m.type}</span>
                        <span>{m.clientId}</span>
                    </div>
                ))}
            </div>
        </main>
    )
}

export default Chat
