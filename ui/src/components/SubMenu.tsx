import * as React from "react";
import { useState } from "react";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import KeyboardArrowRightIcon from "@mui/icons-material/KeyboardArrowRight";
import { List, ListItemButton, ListItemText, Collapse } from "@mui/material";
import { Link, useSidebarState } from "react-admin";

export const SubMenu = (props: SubMenuProps) => {
  const {
    isDropdownOpen = false,
    primaryText,
    leftIcon,
    to,
    selected,
    children,
  } = props;
  const [open] = useSidebarState();
  const [isOpen, setIsOpen] = useState(isDropdownOpen);

  const handleToggle = () => {
    setIsOpen(!isOpen);
  };

  return (
    <React.Fragment>
      <Link to={to}>
        <ListItemButton
          dense
          onClick={handleToggle}
          sx={{
            paddingLeft: "1rem",
            color: "rgba(0, 0, 0, 0.54)",
            backgroundColor: selected ? "rgba(25, 118, 210, 0.08)" : "",
          }}
        >
          {isOpen ? <ExpandMoreIcon /> : leftIcon ?? <KeyboardArrowRightIcon />}
          <ListItemText
            inset
            disableTypography
            primary={primaryText}
            sx={{
              paddingLeft: 2,
              fontSize: "1rem",
              color: "rgba(0, 0, 0, 0.6)",
            }}
          />
        </ListItemButton>
      </Link>
      <Collapse in={isOpen} timeout="auto" unmountOnExit>
        <List
          component="div"
          disablePadding
          sx={
            open
              ? {
                  paddingLeft: "25px",
                  transition:
                    "padding-left 195ms cubic-bezier(0.4, 0, 0.6, 1) 0ms",
                }
              : {
                  paddingLeft: 0,
                  transition:
                    "padding-left 195ms cubic-bezier(0.4, 0, 0.6, 1) 0ms",
                }
          }
        >
          {children}
        </List>
      </Collapse>
    </React.Fragment>
  );
};

export type SubMenuProps = {
  children?: React.ReactNode;
  isDropdownOpen?: boolean;
  leftIcon?: React.ReactElement;
  primaryText?: string;
  to: string;
  selected: boolean;
};

export default SubMenu;
