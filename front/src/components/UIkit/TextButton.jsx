import Button from '@material-ui/core/Button';
import { makeStyles } from '@material-ui/styles';
import { theme } from '../../assets/theme';
import ArrowForwardIosIcon from '@material-ui/icons/ArrowForwardIos';

const useStyles = makeStyles({
  button: {
    color: '#fff',
    fontSize: 20,
    fontWeight: 'bold',
    fontFamily: 'inherit',
    border: '2px solid #fff',
    borderRadius: '35px',
    [theme.breakpoints.down('xs')]: {
      // SP
      fontSize: '16px',
    },
    '& > span > span': {
      // ボタンアイコン
      marginLeft: '5px',
    },
  },
});

const TextButton = (props) => {
  const classes = useStyles();

  return (
    <Button
      className={classes.button}
      variant='outlined'
      onClick={() => props.onClick()}
      endIcon={<ArrowForwardIosIcon />}
    >
      {props.label}
    </Button>
  );
};

export default TextButton;
