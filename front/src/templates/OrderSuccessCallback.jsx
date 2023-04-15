import { useEffect } from 'react';
import { push } from 'connected-react-router';
import { checkCheckoutResult } from '../reducks/users/operations';
import { showMessageAction } from '../reducks/messages/actions';
import { useDispatch } from 'react-redux';
import { useLocation } from 'react-router-dom';

const OrderSuccessCallback = () => {
  const dispatch = useDispatch();

  const search    = useLocation().search;
  const query     = new URLSearchParams(search);
  const sessionId = query.get('session_id');
  const orderId   = query.get('order_id');

  if (!sessionId || !orderId) {
    dispatch(showMessageAction('error', '無効なURLです'));
    dispatch(push('/'));
  } else {
    useEffect(() => {
      dispatch(checkCheckoutResult(sessionId, orderId));
    }, []);
  }

  return (<></>);
};

export default OrderSuccessCallback;
