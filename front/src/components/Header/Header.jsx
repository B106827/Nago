import { useCallback, useState } from 'react';
import { makeStyles } from '@material-ui/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
// 画像を変更した場合フラッシュメッセージの表示位置を変更する
import logo from '../../assets/img/icons/logo.svg';
import { useDispatch } from 'react-redux';
import { push } from 'connected-react-router';
import { HeaderMenus, ClosableDrawer } from './index';

const Header = () => {
  const classes = useStyles();
  const dispatch = useDispatch();

  const [open, setOpen] = useState(false);

  const handleDrawerToggle = useCallback(
    (event) => {
      if (
        event.type === 'keydown' &&
        (event.key === 'Tab' || event.key === 'Shift')
      ) {
        return;
      }
      setOpen(!open);
    },
    [setOpen, open]
  );

  return (
    <div className={classes.root}>
      <AppBar position='fixed' className={classes.menuBar}>
        <Toolbar className={classes.toolBar}>
          <img
            src={logo}
            onClick={() => dispatch(push('/'))}
            className={'cursor-pointer' + ' ' + `${classes.logo}`}
          />
          <div className={classes.iconButtons}>
            <HeaderMenus handleDrawerToggle={handleDrawerToggle} />
          </div>
        </Toolbar>
      </AppBar>
      <ClosableDrawer open={open} onClose={handleDrawerToggle} />
    </div>
  );
};

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
  },
  menuBar: {
    backgroundColor: theme.palette.primary.main,
    boxShadow: 'none',
  },
  toolBar: {
    margin: '0 auto',
    maxWidth: theme.size.window.contentMaxWidth,
    width: '100%',
  },
  iconButtons: {
    margin: '0 0 0 auto',
  },
  logo: {
    width: '128px',
    [theme.breakpoints.up('sm')]: {
      // タブレット以上
      width: '148px',
    },
  },
}));

export default Header;
