/* eslint-disable no-unused-vars */
import { useEffect, useState } from 'react';

const PopupChat = () => {
    const [messages, setMessages] = useState([]);
    const [chatId] = useState(1);
    // const [userId, setUserId] = useState(getUserId());

    useEffect(() => {
        loadMessages();
        const interval = setInterval(() => {
            streamMessages();
        }, 1000);
        return () => clearInterval(interval);
    }, []);

    const getUserId = () => {
        const value = `; ${document.cookie}`;
        const parts = value.split(`; user-id=`);
        if (parts.length === 2) return parseInt(parts.pop().split(';').shift(), 10);
        return null;
    };

    const loadMessages = () => {
        fetch(`http://localhost:8081/api/user/get-messages/${chatId}`, { method: "GET" })
            .then(response => response.json())
            .then(data => {
                setMessages(data.messages);
            });
    };

    const streamMessages = () => {
        fetch(`http://localhost:8081/api/user/get-messages/${chatId}`, { method: "GET" })
            .then(response => response.json())
            .then(data => {
                setMessages(prevMessages => {
                    const messageIds = prevMessages.map(msg => msg.id);
                    const newMessages = data.messages.filter(msg => !messageIds.includes(msg.id));
                    return [...prevMessages, ...newMessages];
                });
            });
    };

    const sendMessage = (content) => {
        fetch(`http://localhost:8081/api/user/send-message`, {
            method: "POST",
            headers: {
                'Access-Control-Allow-Origin': '*',
                'Content-Type': 'application/json;charset=utf-8',
            },
            body: JSON.stringify({
                "message": {
                    "id": 0,
                    "id_author": userId,
                    "id_chat": chatId,
                    "content": content,
                    "datetime": new Date().toISOString()
                }
            })
        });
    };

    const handleSendMessage = () => {
        const input = document.querySelector(".popup-chat__textbox > input");
        const content = input.value;
        input.value = '';
        sendMessage(content);
    };

    return (
        <div className="popup-chat">
            <div className="popup-chat__header">
                <div className="friend" style={{ backgroundColor: '#2f2f2f' }}>
                    <img className="friend__avatar" src="/images/Ellipse.png" alt="Avatar" />
                    <div className="friend__data">
                        <p className="friend__username">Ivan</p>
                        <p className="friend__status">online</p>
                    </div>
                </div>
            </div>
            <div className="popup-chat__content">
                {messages.map(message => (
                    <div key={message.id} className={`message ${message.id_author === userId ? 'my-message' : 'friend-message'}`}>
                        {message.content}
                    </div>
                ))}
            </div>
            <div className="popup-chat__textbox">
                <input type="text" placeholder="Введите текст" />
                <a onClick={handleSendMessage} className="friend__chat">
                    <img src="/images/comment.svg" alt="Send" />
                </a>
            </div>
        </div>
    );
};

export default PopupChat;