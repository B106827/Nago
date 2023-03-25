import { useEffect } from 'react';
import { useDispatch } from 'react-redux';
import List from '@material-ui/core/List';
import { makeStyles } from '@material-ui/styles';
import { fetchMyInfo } from '../reducks/users/operations';

const useStyles = makeStyles((theme) => ({
  myProfile: {
    background: theme.palette.grey['100'],
    margin: '0 auto',
    padding: 32,
    [theme.breakpoints.down('md')]: {
      width: '100%',
    },
    [theme.breakpoints.up('md')]: {
      width: 768,
    },
  },
}));

const UserMyPage = () => {
  const classes = useStyles();
  const dispatch = useDispatch();

  useEffect(() => {
    dispatch(fetchMyInfo());
  }, []);

  return (
    <section className='c-section-wrapin'>
      <List className={classes.myProfile}>
      </List>
    </section>
  );
};

export default UserMyPage;
