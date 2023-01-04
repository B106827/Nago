import Router from './Router';
import './assets/reset.css';
import './assets/style.css';
import { Header } from './components/Header';
import { Message } from './components/UIkit';

const App = () => {
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
