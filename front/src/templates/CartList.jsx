import { useCallback } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import List from '@material-ui/core/List';
import { getMyCartList } from '../reducks/users/selectors';
import { CartListItem } from '../components/Products';
import { PrimaryButton } from '../components/UIkit';
import { push } from 'connected-react-router';
import { makeStyles } from '@material-ui/styles';

const CartList = () => {
  const classes = useStyles();
  const dispatch = useDispatch();
  const selector = useSelector((state) => state);
  const cartList = getMyCartList(selector);

  const goToOrder = useCallback(() => {
    dispatch(push('/order/confirm'));
  }, [dispatch]);

  const backToHome = useCallback(() => {
    dispatch(push('/'));
  }, [dispatch]);

  console.log(cartList);
  return (
    <section>
      <div className={classes.topSection}>
        <h2 className={classes.topSectionTitle}>カート</h2>
        <List className={classes.topSectionCartList}>
          {cartList.length > 0 &&
            cartList.map((cart) => (
              <CartListItem key={cart.id} cart={cart} />
            ))
          }
        </List>
        <div className='module-spacer--medium' />
        <div className='p-grid__column' />
        <div className={classes.topSectionButton} >
          <PrimaryButton label={'レジへ進む'} onClick={goToOrder} />
          <div className='module-spacer--extra-extra-small' />
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
  topSectionCartList: {
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

export default CartList;
