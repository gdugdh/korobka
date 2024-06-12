import Sidebar from '../widgets/Sidebar';
import MainContent from '../widgets/MainContent';
import FriendsList from '../widgets/FriendsList';

const HomePage = () => (
        <div className="content__container container">
            <Sidebar />
            <MainContent />
            <FriendsList />
        </div>
);

export default HomePage;