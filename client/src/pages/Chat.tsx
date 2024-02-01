import { FormEvent, useCallback, useEffect, useRef, useState } from 'react'

const Chat = () => {
    const socketRef = useRef(new WebSocket("ws://localhost:8080/chat"))
    const [massage, setMassage] = useState("");


    const handelSend = (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        if (massage) {
            console.log("I am Sending Stuff")
            const payload = JSON.stringify({ type: 'chat message', massage })
            socketRef.current.send(payload);
            setMassage('');
        }
        else console.log(" I Dose Not Do anyThing")
    }

    const handler = useCallback((x: MessageEvent<unknown>) => {
        console.log(x)
    }, [])

    useEffect(() => {
        socketRef.current.addEventListener("message", handler);
        return () => {
            // eslint-disable-next-line react-hooks/exhaustive-deps
            socketRef.current.removeEventListener("message", handler)
        }
    }, [handler])

    return (
        <main>
            <form onSubmit={(e) => handelSend(e)} className="flex flex-col gap-2">
                <label htmlFor="send-massage-input">Send Massage</label>
                <input type="text" id="send-massage-input" value={massage} onChange={(e) => setMassage(e.target.value)} placeholder="..." />
                <button type="submit">Send</button>
            </form>
        </main>
    )
}

export default Chat
