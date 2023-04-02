import TextField from '@material-ui/core/TextField';
import { makeStyles } from '@material-ui/styles';

const TextInput = (props) => {
  const classes = useStyles();
  return (
    <TextField
      fullWidth={props.fullWidth}
      label={props.label}
      placeholder={props.placeholder}
      helperText={props.helperText}
      margin='dense'
      multiline={props.multiline}
      required={props.required}
      rows={props.rows}
      value={props.value}
      type={props.type}
      onChange={props.onChange}
      variant={props.variant}
      InputProps={props.InputProps}
      InputLabelProps={{
        shrink: true,
      }}
      className={classes.root}
    />
  );
};

const useStyles = makeStyles({
  root: {
    '& .MuiInputLabel-root': {
      fontSize: '1.3rem'
    },
    '& .MuiInput-input': {
      paddingTop: 10
    },
  },
});

export default TextInput;
