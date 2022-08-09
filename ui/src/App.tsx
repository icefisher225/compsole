import {
  AppBar,
  Avatar,
  Box,
  Button,
  Divider,
  IconButton,
  Menu,
  MenuItem,
  Toolbar,
  Tooltip,
  Typography,
} from "@mui/material";
import React, { useEffect, useState } from "react";
import { Link, Outlet, useNavigate, useLocation } from "react-router-dom";
import { Role, useGetCurrentUserQuery, User } from "./api/generated/graphql";
import { Loading } from "./pages/loading";
import { UserContext } from "./user-context";
import SettingsIcon from "@mui/icons-material/Settings";
import PowerSettingsNewIcon from "@mui/icons-material/PowerSettingsNew";
import { Logout } from "./api";

import Logo from "./res/logo_200.png";
import { useSnackbar } from "notistack";

function App() {
  const {
    data: currentUser,
    loading: currentUserLoading,
    error: currentUserError,
  } = useGetCurrentUserQuery();
  let navigate = useNavigate();
  let location = useLocation();
  const { enqueueSnackbar } = useSnackbar();
  let [user, setUser] = useState<User | null | undefined>();
  let [userMenuOpen, setUserMenuOpen] = useState<boolean>(false);

  let handleOpenUserMenu = () => setUserMenuOpen(true);
  let handleCloseUserMenu = () => setUserMenuOpen(false);

  // Ensure the user cannot access protected App without authentication
  useEffect(() => {
    if (!currentUserLoading && (currentUserError || !currentUser))
      navigate("/auth/signin", {
        state: {
          from: location,
        },
      });
    else if (!currentUserLoading && !currentUserError && currentUser)
      setUser(currentUser.me);
  }, [currentUser, currentUserLoading, currentUserError, navigate, location]);

  const handleSignOut = () => {
    Logout().then(
      () => {
        enqueueSnackbar("Logged out successfully", {
          variant: "success",
        });
        navigate("/auth/signin");
      },
      (err) => {
        enqueueSnackbar(`Failed to logout: ${err}`, {
          variant: "error",
        });
      }
    );
  };

  return !currentUserLoading && user ? (
    <UserContext.Provider value={user}>
      <Box sx={{ flexGrow: 1 }}>
        <AppBar position="static">
          <Toolbar>
            <Link
              to="/"
              style={{
                flexGrow: 1,
                textDecoration: "none",
                color: "inherit",
                display: "flex",
                alignItems: "center",
              }}
            >
              <img
                src={Logo}
                alt="Logo"
                style={{ height: "2.5rem", marginRight: "0.5rem" }}
              />
              <Typography variant="h6">Compsole</Typography>
            </Link>
            <Box
              sx={{
                display: "flex",
                flex: 1,
                alignItems: "center",
                justifyContent: "flex-end",
                "& hr": {
                  mx: 0.5,
                },
              }}
            >
              <Button href="/" color="inherit">
                Dashboard
              </Button>
              {user && user.Role === Role.Admin && (
                <>
                  <Button href="/admin" color="inherit">
                    Admin
                  </Button>
                  <Button href="/admin/ingest" color="inherit">
                    Ingest VMs
                  </Button>
                </>
              )}
              <Divider orientation="vertical" variant="middle" flexItem />
              {/*
              <Typography
                variant="subtitle2"
                sx={{
                  ml: 1,
                }}
              >
                Hello, {user.FirstName} {user.LastName}
              </Typography> */}
              <Box sx={{ flexGrow: 0 }}>
                <Tooltip title="Open settings">
                  <IconButton onClick={handleOpenUserMenu} sx={{ p: 0, ml: 1 }}>
                    <Avatar
                      alt={`${user.FirstName} ${user.LastName}`}
                      sx={{
                        bgcolor: "#F7B374",
                      }}
                    >
                      {user.FirstName.at(0)}
                      {user.LastName.at(0)}
                    </Avatar>
                  </IconButton>
                </Tooltip>
                <Menu
                  sx={{ mt: "45px" }}
                  id="menu-appbar"
                  anchorOrigin={{
                    vertical: "top",
                    horizontal: "right",
                  }}
                  keepMounted
                  transformOrigin={{
                    vertical: "top",
                    horizontal: "right",
                  }}
                  open={userMenuOpen}
                  onClose={handleCloseUserMenu}
                >
                  <MenuItem onClick={handleCloseUserMenu}>
                    <SettingsIcon sx={{ mr: 1 }} />
                    <Typography textAlign="left">Account Settings</Typography>
                  </MenuItem>
                  <MenuItem onClick={handleSignOut}>
                    <PowerSettingsNewIcon sx={{ mr: 1 }} />
                    <Typography textAlign="left">Sign Out</Typography>
                  </MenuItem>
                </Menu>
              </Box>
            </Box>
          </Toolbar>
        </AppBar>
      </Box>
      <Outlet />
    </UserContext.Provider>
  ) : (
    <Loading />
  );
}

export default App;
