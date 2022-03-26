import { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { getIsLogedIn } from './reducks/users/selectors';
import { listenAuth } from './reducks/users/operations';

const Auth = ({ children }) => {
  const dispatch = useDispatch();
  const selector = useSelector((state) => state);
  const isLogedIn = getIsLogedIn(selector);

  useEffect(() => {
    if (!isLogedIn) {
      dispatch(listenAuth());
    }
  }, [dispatch, isLogedIn]);

  if (!isLogedIn) {
    return <></>;
  } else {
    return children;
  }
};

export default Auth;
