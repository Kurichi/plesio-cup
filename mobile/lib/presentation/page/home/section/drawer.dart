import 'package:firebase_auth/firebase_auth.dart';
import 'package:flutter/material.dart';
import 'package:kiikuten/presentation/page/settings/settings_screen.dart';

class KiikutenDrawer extends StatelessWidget {
  const KiikutenDrawer({super.key});

  @override
  Widget build(BuildContext context) {
    return Drawer(
      child: ListView(
        children: [
          ListTile(
            title: const Text('木育展'),
            onTap: () {
              Navigator.pop(context);
            },
          ),
          ListTile(
            title: const Text('設定'),
            onTap: () {
              Navigator.pop(context);
              Navigator.of(context).push(
                MaterialPageRoute(
                  builder: (context) => const SettingsScreen(),
                ),
              );
            },
          ),
          ListTile(
            title: const Text('ログアウト'),
            onTap: () {
              Navigator.pop(context);
              FirebaseAuth.instance.signOut();
            },
          ),
        ],
      ),
    );
  }
}
