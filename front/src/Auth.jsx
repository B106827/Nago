import { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { getIsLogedIn } from './reducks/users/selectors';
import { showMessageAction } from './reducks/messages/actions';

const Auth = ({ children }) => {
  const dispatch = useDispatch();
  const selector = useSelector((state) => state);
  const isLogedIn = getIsLogedIn(selector);

  useEffect(() => {
    if (!isLogedIn) {
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
