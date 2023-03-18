import { useEffect, useState }from 'react';
import IconButton from '@material-ui/core/IconButton';
import Badge from '@material-ui/core/Badge';
import ShoppingCartOutlinedIcon from '@material-ui/icons/ShoppingCartOutlined';
import MenuIcon from '@material-ui/icons/Menu';
import { fetchMyCartList } from '../../reducks/users/operations';
import { getMyCartList, getIsLogedIn } from '../../reducks/users/selectors';
import { useDispatch, useSelector } from 'react-redux';
import { push } from 'connected-react-router';
import { makeStyles } from '@material-ui/styles';

const HeaderMenus = (props) => {
  const classes = useStyles();
  const dispatch = useDispatch();
  const selector = useSelector((state) => state);
  const isLogedIn = getIsLogedIn(selector);

  useEffect(() => {
    if (isLogedIn) {
      dispatch(fetchMyCartList());
    }
  }, []);

  const [cartList, setCartList] = useState([]);

  const storeMyCartList = getMyCartList(selector);
  useEffect(() => {
    if (storeMyCartList) {
      setCartList(storeMyCartList);
    }
  }, [storeMyCartList]);

  return (
    <>
      <IconButton onClick={() => dispatch(push('/cart'))}>
        <Badge
          badgeContent={cartList && cartList.length}
          max={99}
          classes={{ badge: classes.badge }}
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

const useStyles = makeStyles((theme) => ({
  badge: {
    color: theme.palette.primary.black,
    backgroundColor: theme.palette.primary.light
  },
  icon: {
    color: theme.palette.primary.light,
    fontSize: 30,
  },
}));

export default HeaderMenus;
