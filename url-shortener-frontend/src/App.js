import './App.css';
import './styles/styles.css';
import HomePage from './pages/home.tsx'
import NotFound from './pages/pageNotFound.tsx'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import LoginPage from './pages/login.tsx';


function App() {
  return (
    <div className="app">
      <Router>
        <Routes>
        <Route
            exact
            path='/auth'
            element={<LoginPage />}
          ></Route>
          <Route
            exact
            path='/'
            element={<HomePage />}
          ></Route>

          <Route path="*" element={<NotFound />} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
