using System;
using System.Diagnostics;

namespace Thetis
{
    // W5TSU: launches and supervises the external FreeDV Reporter helper
    // process (thetisctl or its thetis-ai-skill successor -- see
    // docs/superpowers/specs/2026-08-26-freedv-tab-and-self-reporting-design.md).
    // Follows Dumpcap.cs's existing exact pattern in this codebase: a
    // static class, Process.Start with CreateNoWindow, PID-tracked kill,
    // no stdout/stderr redirection (matching this project's only existing
    // precedent for supervising a long-running external process).
    public static class FreeDVReporterHelper
    {
        private static int m_nProcessID = -1;
        private const string PROCESS_NAME_NO_EXT = "thetisctl";

        public static bool IsRunning
        {
            get
            {
                if (m_nProcessID == -1) return false;

                bool bRet = false;
                Process[] proc = Process.GetProcessesByName(PROCESS_NAME_NO_EXT);
                foreach (Process p in proc)
                {
                    if (p.Id == m_nProcessID)
                    {
                        bRet = true;
                        break;
                    }
                }
                return bRet;
            }
        }

        // Starts (or restarts, if already running with different
        // arguments) the helper process. helperPath empty means resolve
        // "thetisctl" via PATH. Returns true if a process is running
        // afterward (either newly started or an unchanged prior instance
        // — callers pass the same arguments they'd want running now, so a
        // restart-on-every-call keeps this simple).
        public static bool EnsureRunning(string helperPath, string arguments)
        {
            Stop();

            try
            {
                string fileName = string.IsNullOrEmpty(helperPath) ? PROCESS_NAME_NO_EXT : helperPath;

                using (Process myProcess = new Process())
                {
                    myProcess.StartInfo.UseShellExecute = false;
                    myProcess.StartInfo.FileName = fileName;
                    myProcess.StartInfo.Arguments = arguments;
                    myProcess.StartInfo.CreateNoWindow = true;
                    myProcess.Start();
                    m_nProcessID = myProcess.Id;
                }
                return true;
            }
            catch
            {
                m_nProcessID = -1;
                return false;
            }
        }

        public static void Stop()
        {
            if (!IsRunning) return;

            Process[] proc = Process.GetProcessesByName(PROCESS_NAME_NO_EXT);
            foreach (Process p in proc)
            {
                if (p.Id == m_nProcessID)
                {
                    try
                    {
                        p.Kill();
                        m_nProcessID = -1;
                    }
                    catch { }
                    break;
                }
            }
        }
    }
}
