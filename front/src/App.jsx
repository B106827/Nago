import Router from './Router';
import { useDispatch } from 'react-redux';
import './assets/reset.css';
import './assets/style.css';
import { Header } from './components/Header';
import { Message } from './components/UIkit';
import { listenAuth } from './reducks/users/operations';

const App = () => {
  // ページリフレッシュされてもログイン状態を維持できるように
  const dispatch = useDispatch();
  dispatch(listenAuth(false));

  return (
    <>
      <Header />
      <Message>
        <main className='c-main'>
          <Router />
        </main>
      </Message>
    </>
  );
};

export default App;
