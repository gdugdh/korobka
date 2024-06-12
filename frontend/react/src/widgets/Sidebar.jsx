/* eslint-disable no-unused-vars */
import { useState, useEffect } from 'react';
import addIcon from '/images/add.svg';
import ballIcon from '/images/ball.svg';
import volleyballIcon from '/images/volleyball-ball.svg';
import basketballIcon from '/images/basketball-ball.svg';
import tennisIcon from '/images/tennis-ball.svg';
import arrowIcon from '/images/arrow.svg';

const Sidebar = () => {
    const [disciplines, setDisciplines] = useState([
        { id: 1, name: 'Футбол', unactive: ballIcon, active: 'image/ball-active.svg' },
        { id: 2, name: 'Волейбол', unactive: volleyballIcon, active: 'image/volleyball-ball-active.svg' },
        { id: 3, name: 'Баскетбол', unactive: basketballIcon, active: 'image/basketball-ball-active.svg' },
        { id: 4, name: 'Теннис', unactive: tennisIcon, active: 'image/tennis-ball-active.svg' }
    ]);
    const [choiceDisciplineId, setChoiceDisciplineId] = useState(1);
    const [matches, setMatches] = useState([]);

    useEffect(() => {
        choiceDiscipline(choiceDisciplineId);
    }, [choiceDisciplineId]);

    const changeActiveDisciplineItem = (disciplineId) => {
        setDisciplines(prevDisciplines =>
            prevDisciplines.map(discipline =>
                discipline.id === disciplineId
                    ? { ...discipline, active: true }
                    : { ...discipline, active: false }
            )
        );
    };

    const choiceDiscipline = (disciplineId) => {
        changeActiveDisciplineItem(disciplineId);

        fetch(`http://localhost:8081/api/match/${disciplineId}/active`, { method: "GET" })
            .then(response => response.json())
            .then(data => {
                setMatches(data.matches);
            });
    };

    return (
        <div className="sidebar">
            <div className="content-block">
                <h3>Дисциплины</h3>
                <div className="discipline-list">
                    <a className="discipline-item">
                        <img src={addIcon} alt="Add" />
                        <p className="discipline-item__text">Создать игру</p>
                    </a>
                    {disciplines.map(discipline => (
                        <a key={discipline.id} className={`discipline-item ${discipline.id === choiceDisciplineId ? 'active' : ''}`} onClick={() => setChoiceDisciplineId(discipline.id)}>
                            <img className="discipline-item__icon" src={discipline.unactive} alt={discipline.name} />
                            <p className="discipline-item__text">{discipline.name}</p>
                            <img src={arrowIcon} className="discipline-item__arrow" alt="Arrow" />
                        </a>
                    ))}
                </div>
            </div>
        </div>
    );
};

export default Sidebar;
