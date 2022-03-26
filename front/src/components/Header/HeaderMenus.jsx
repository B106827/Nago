import IconButton from '@material-ui/core/IconButton';
import Badge from '@material-ui/core/Badge';
import ShoppingCartOutlinedIcon from '@material-ui/icons/ShoppingCartOutlined';
import MenuIcon from '@material-ui/icons/Menu';
import { getProductsInCart } from '../../reducks/users/selectors';
import { useDispatch, useSelector } from 'react-redux';
import { push } from 'connected-react-router';
import { makeStyles } from '@material-ui/styles';

const HeaderMenus = (props) => {
  const classes = useStyles();
  const dispatch = useDispatch();
  const selector = useSelector((state) => state);
  const productsInCart = getProductsInCart(selector);

  //useEffect(() => {
  //  const unsubscribe = db.collection('users').doc(uid).collection('cart')
  //    .onSnapshot(snapshots => {
  //      snapshots.docChanges().forEach(change => {
  //        const product = change.doc.data();
  //        const changeType = change.type;

  //        switch (changeType) {
  //          case 'added':
  //            productsInCart.push(product);
  //            break;
  //          case 'modified':
  //            const index = productsInCart.findIndex(product => product.cartId === change.doc.id);
  //            productsInCart[index] = product;
  //            break;
  //          case 'removed':
  //            productsInCart = productsInCart.filter(product => product.cartId !== change.doc.id);
  //            break;
  //          default:
  //            break;
  //        }
  //      })
  //      dispatch(fetchProductsInCart(productsInCart));
  //    })

  //  return () => unsubscribe();
  //}, []);

  return (
    <>
      <IconButton onClick={() => dispatch(push('/cart'))}>
        <Badge
          badgeContent={productsInCart && productsInCart.length}
          color='secondary'
        >
          <ShoppingCartOutlinedIcon className={classes.icon} />
        </Badge>
      </IconButton>
      <IconButton onClick={(event) => props.handleDrawerToggle(event)}>
        <MenuIcon className={classes.icon} />
      </IconButton>
    </>
  );
};

const useStyles = makeStyles(() => ({
  icon: {
    color: '#fff',
    fontSize: 30,
  },
}));

export default HeaderMenus;
