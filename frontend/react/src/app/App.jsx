import { Routes, Route } from 'react-router-dom';
import HomePage from '../pages/HomePage';
import SearchPage from '../pages/SearchPage';
import Header from '../widgets/Header';
import PopupChat from '../widgets/PopupChat';
import './styles/reset.css';
import './styles/style.css';

const App = () => (
  <div className='content'>
        <Header/>
        <Routes>
            <Route path="/" element={<HomePage />} />
            <Route path="/Search" element={<SearchPage />} />
        </Routes>
    <PopupChat />
    </div>
);

export default App;
