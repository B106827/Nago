import { useEffect, useCallback, useState } from 'react';
import List from '@material-ui/core/List';
import { useDispatch, useSelector } from 'react-redux';
import { makeStyles } from '@material-ui/styles';
import { getOrderHistoryList } from '../reducks/users/selectors';
import { fetchOrderHistory } from '../reducks/users/operations';
import { OrderHistoryItem } from '../components/Products';
import { push } from 'connected-react-router';
import { PrimaryButton } from '../components/UIkit';

const OrderHistory = () => {
  const classes  = useStyles();
  const dispatch = useDispatch();
  const selector = useSelector((state) => state);

  useEffect(() => {
    dispatch(fetchOrderHistory());
  }, []);
  const storeOrderHistoryList = getOrderHistoryList(selector);

  const [orderHistoryList, setOrderHistoryList] = useState(null);

  useEffect(() => {
    if (Array.isArray(storeOrderHistoryList) && storeOrderHistoryList.length > 0) {
      setOrderHistoryList(storeOrderHistoryList);
    }
  }, [storeOrderHistoryList]);

  const backToHome = useCallback(() => {
    dispatch(push('/'));
  }, [dispatch]);

  return (
    <section>
      <div className={classes.topSection}>
        <h2 className={classes.topSectionTitle}>注文履歴</h2>
        <List className={classes.topSectionOrderHistoryList}>
          {orderHistoryList && orderHistoryList.length > 0 ? (
            orderHistoryList.map((orderHistory) => (
              <OrderHistoryItem key={orderHistory.id} order={orderHistory} />
            ))
          ) : (
            <p>購入履歴はありません</p>
          )}
        </List>
        <div className='module-spacer--medium' />
        <div className='p-grid__column' />
        <div className={classes.topSectionButton} >
          <PrimaryButton label={'ショッピングを続ける'} onClick={backToHome} />
        </div>
      </div>
    </section>
  );
};

const useStyles = makeStyles((theme) => ({
  // セクション1
  topSection: {
    padding: '40px 20px 0',
    [theme.breakpoints.up('md')]: {
      // PC
      padding: '60px 40px',
    },
  },
  topSectionTitle: {
    textAlign: 'center',
    fontWeight: 'bold',
    marginBottom: '32px',
    [theme.breakpoints.down('xs')]: {
      // SP
      fontSize: '20px',
    },
    [theme.breakpoints.up('sm')]: {
      // タブレット
      fontSize: '24px',
    },
    [theme.breakpoints.up('md')]: {
      // PC
      fontSize: '28px',
    },
  },
  topSectionOrderHistoryList: {
    margin: '0 auto',
    maxWidth: 512,
    width: '100%',
  },
  topSectionButton: {
    margin: '0 auto',
    maxWidth: 512,
    width: '100%',
    textAlign: 'center',
  },
}));


export default OrderHistory;
