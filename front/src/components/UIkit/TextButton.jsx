import Button from '@material-ui/core/Button';
import { makeStyles } from '@material-ui/styles';
import ArrowForwardIosIcon from '@material-ui/icons/ArrowForwardIos';

const TextButton = (props) => {
  const classes = useStyles(props);

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

// スタイル
const useStyles = makeStyles((theme) => ({
  button: (props) => ({
    color:  (props.style?.color || theme.palette.primary.black),
    backgroundColor: (props.style?.bgColor || theme.palette.primary.light),
    fontSize: 16,
    fontWeight: 'bold',
    fontFamily: 'inherit',
    border: '2px solid #fff',
    borderRadius: '35px',
    '& > span > span': {
      // ボタンアイコン
      marginLeft: '5px',
    },
  }),
}));


export default TextButton;
