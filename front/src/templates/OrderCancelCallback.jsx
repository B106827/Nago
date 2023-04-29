import { useEffect } from 'react';
import { useDispatch } from 'react-redux';
import { showMessageAction } from '../reducks/messages/actions';
import { cancelCheckout } from '../reducks/users/operations';
import { useLocation } from 'react-router-dom';
import { push } from 'connected-react-router';

const OrderSuccessCallback = () => {
  const dispatch = useDispatch();
  const search   = useLocation().search;
  const query    = new URLSearchParams(search);
  const orderId  = query.get("order_id")

  if (!orderId) {
    dispatch(showMessageAction('error', '無効なURLです'));
    dispatch(push('/'));
  } else {
    useEffect(() => {
      dispatch(cancelCheckout(orderId));
    }, []);
  }

  return (<></>);
};

export default OrderSuccessCallback;
