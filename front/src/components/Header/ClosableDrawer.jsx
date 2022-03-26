import { useDispatch, useSelector } from 'react-redux';
import Divider from '@material-ui/core/Divider';
import Drawer from '@material-ui/core/Drawer';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import { makeStyles } from '@material-ui/styles';
import { push } from 'connected-react-router';
import { getIsLogedIn } from '../../reducks/users/selectors';
import config from '../../config/base';
import { logout } from '../../reducks/users/operations';

const ClosableDrawer = (props) => {
  const classes = useStyles();
  const { container } = props;
  const dispatch = useDispatch();
  const selector = useSelector((state) => state);

  const selectMenu = (event, path) => {
    if (path === 'logout') {
      dispatch(logout());
      props.onClose(event, false);
      return;
    }
    dispatch(push(path));
    props.onClose(event, false);
  };

  const isLogedIn = getIsLogedIn(selector);

  const menus = isLogedIn
    ? [...config.menus.authenticated]
    : [...config.menus.noAuthenticated];

  return (
    <nav className={classes.drawer}>
      <Drawer
        container={container}
        variant='temporary'
        anchor='right'
        open={props.open}
        onClose={(e) => props.onClose(e, false)}
        classes={{ paper: classes.drawerPaper }}
        ModalProps={{ keepMounted: true }}
      >
        <div
          onClose={(e) => props.onClose(e)}
          onKeyDown={(e) => props.onClose(e)}
        >
          <Divider />
          <List>
            {menus.map((menu, index) => (
              <ListItem
                button
                key={index}
                onClick={(e) => selectMenu(e, menu.value)}
              >
                <ListItemIcon>{menu.icon}</ListItemIcon>
                <ListItemText primary={menu.label} />
              </ListItem>
            ))}
          </List>
          <Divider />
        </div>
      </Drawer>
    </nav>
  );
};

const useStyles = makeStyles((theme) => ({
  drawer: {
    [theme.breakpoints.up('sm')]: {
      flexShrink: 0,
      width: 256,
    },
  },
  toolbar: theme.mixins.toolbar,
  drawerPaper: {
    width: 256,
  },
  searchField: {
    alignItems: 'center',
    display: 'flex',
    marginLeft: 32,
  },
}));

export default ClosableDrawer;
