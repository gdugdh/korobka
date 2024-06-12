import { useEffect, useState } from 'react';

const FriendsList = () => {
    const [friends, setFriends] = useState([]);

    useEffect(() => {
        fetch(`http://localhost:8081/api/user/${getUserId()}/friends`, { method: "GET" })
            .then(response => response.json())
            .then(data => {
                setFriends(data.users);
            });
    }, []);

    const getUserId = () => {
        const value = `; ${document.cookie}`;
        const parts = value.split(`; user-id=`);
        if (parts.length === 2) return parseInt(parts.pop().split(';').shift(), 10);
        return null;
    };

    return (
        <div className="content-block">
            <h3>Друзья</h3>
            <div className="friend-list">
                {friends.length ? friends.map(friend => (
                    <div key={friend.id} className="friend">
                        <img className="friend__avatar" src="image/Ellipse.png" alt="Avatar" />
                        <div className="friend__data">
                            <p className="friend__username">{friend.username}</p>
                            <p className="friend__status">{friend.status}</p>
                        </div>
                        <div className="friend__chat">
                            <svg className="friend__chat-image" width="21" height="21" viewBox="0 0 21 21" fill="none" xmlns="http://www.w3.org/2000/svg">
                                <path d="M0.5 20.5V2.72222C0.5 1.49492 1.49492 0.5 2.72222 0.5H18.2778C19.5051 0.5 20.5
                                 1.49492 20.5 2.72222V13.8333C20.5 15.0606 19.5051 16.0556 18.2778 16.0556H7.16667C6.6857 
                                 16.0547 6.21757 16.2107 5.83333 16.5L0.5 20.5ZM2.72222 2.72222V16.0556L5.09333 14.2778C5.47738
                                  13.9881 5.94564 13.832 6.42667 13.8333H18.2778V2.72222H2.72222Z" fill="#C0C0C2"/>
                            </svg>
                        </div>
                    </div>
                )) : <p>Нет друзей</p>}
            </div>
        </div>
    );
};

export default FriendsList;
