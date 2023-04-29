import { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { getIsLogedIn } from './reducks/users/selectors';
import { logoutAction } from './reducks/users/actions';
import { showMessageAction } from './reducks/messages/actions';
import { push } from 'connected-react-router';

const Auth = ({ children }) => {
  const dispatch = useDispatch();
  const selector = useSelector((state) => state);
  const isLogedIn = getIsLogedIn(selector);

  useEffect(() => {
    if (!isLogedIn) {
      dispatch(logoutAction());
      dispatch(push('/'));
      dispatch(showMessageAction('error', 'ログインしてください'));
    }
  }, []);

  if (!isLogedIn) {
    return <></>;
  } else {
    return children;
  }
};

export default Auth;
