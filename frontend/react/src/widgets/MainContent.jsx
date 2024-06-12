import { useEffect, useState } from 'react';

const MainContent = () => {
    const [matches, setMatches] = useState([]);

    useEffect(() => {
        // Загрузить активные игры при монтировании компонента
        fetch('http://localhost:8081/api/match/1/active', { method: "GET" })
            .then(response => response.json())
            .then(data => setMatches(data.matches));
    }, []);

    const getStatusString = (match) => {
        if (match.teams[0].club.id === 0 || match.teams[1].club.id === 0) {
            return "Ведется набор игроков";
        }
        return "Ожидается";
    };

    const getTeamHTML = (team) => {
        if (team.club.id !== 0) {
            return (
                <div className="team">
                    <p className="team__name">{team.club.name}</p>
                    <p className="team__description">{team.club.description}</p>
                </div>
            );
        } else {
            return (
                <div className="team">
                    <div className="team__container-count-user">
                        <img src="/images/friends.svg" alt="Friends" />
                        <p>{team.curCountPlayer}/{team.countPlayer}</p>
                    </div>
                    <button className="team__button">
                        Присоединиться
                    </button>
                </div>
            );
        }
    };

    return (
        <div className="mainbar">
            <div className="content-block">
                <h3>Активные игры</h3>
                <div className="match-list">
                    {matches.length ? matches.map(match => (
                        <div key={match.id} className="match-item">
                            <p className="match-item__date">{match.datetimeStart}</p>
                            <p className="match-item__decription">{match.description}</p>
                            <p>{getStatusString(match)}</p>
                            <div className="teams">
                                {getTeamHTML(match.teams[0])}
                                <div className="match-discipline">
                                    <svg width="46" height="45" viewBox="0 0 27 26" fill="none" xmlns="http://www.w3.org/2000/svg">
                                        <path d="M13.3905 8.2L18.0366 11.5167M13.3905 8.2L8.74452 11.5167M13.3905 8.2V4.6M18.0366 11.5167L16.262 16.8833M18.0366 11.5167L21.3288 10M8.74452 11.5167L10.5191 16.8833M8.74452 11.5167L5.4522 10M13.3905 4.6L9.26789 1.70091M13.3905 4.6L17.5132 1.70091M16.262 16.8833H10.5191M16.262 16.8833L18.2756 19.6M21.3288 10L25.5429 14.2M21.3288 10L22.5501 5.06253M10.5191 16.8833L8.50539 19.6M18.2756 19.6L23.9694 19M18.2756 19.6L15.2224 24.866M8.50539 19.6L2.81158 19M8.50539 19.6L11.5586 24.866M5.4522 10L1.23804 14.2M5.4522 10L4.23093 5.06253M25.6033 13C25.6033 19.6274 20.1354 25 13.3905 25C6.64557 25 1.17773 19.6274 1.17773 13C1.17773 6.37258 6.64557 1 13.3905 1C20.1354 1 25.6033 6.37258 25.6033 13Z" stroke="white" strokeLinecap="round" strokeLinejoin="round"/>
                                    </svg>
                                </div>
                                {getTeamHTML(match.teams[1])}
                            </div>
                            <p>{match.location.address}</p>
                        </div>
                    )) : <p style={{ padding: '5px' }}>Активных матчей пока нет</p>}
                </div>
            </div>
        </div>
    );
};

export default MainContent;