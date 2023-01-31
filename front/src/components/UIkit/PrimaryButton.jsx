import Button from '@material-ui/core/Button';
import { makeStyles } from '@material-ui/styles';

const PrimaryButton = (props) => {
  const classes = useStyles();
  const addStyle = props.addStyle;

  return (
    <Button
      className={classes.button}
      style={{ ...addStyle }}
      variant='contained'
      onClick={() => props.onClick()}
    >
      {props.label}
    </Button>
  );
};

const useStyles = makeStyles({
  button: {
    backgroundColor: '#4dd0e1',
    color: '#fff',
    fontSize: '16px',
    fontFamily: 'inherit',
    height: 48,
    width: 256,
  },
});

export default PrimaryButton;
