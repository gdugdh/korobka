import logo from '/images/logo.png';
import commentIcon from '/images/comment.svg';
import staticIcon from '/images/Static.svg';
import profileIcon from '/images/Ellipse.png';

const Header = () => (
    <div className="header">
        <div className="container header__container">
            <div className="header__section">
                <div className="header__logo">
                    <img src={logo} alt="Logo" />
                </div>
            </div>
            <div className="header__section">
                <a href="#" className="header__link">Команды</a>
                <a href="#" className="header__link">Залы</a>
                <a href="#" className="header__link">Контакты</a>
            </div>
            <div className="header__section" style={{ width: '250px' }}></div>
            <div className="header__section header__user-section">
                <a className="header__link">
                    <img src={commentIcon} alt="Comments" />
                </a>
                <a className="header__link">
                    <img src={staticIcon} alt="Stats" />
                </a>
                <a className="header__link">
                    <img src={profileIcon} alt="Profile" />
                </a>
            </div>
        </div>
    </div>
);

export default Header;
